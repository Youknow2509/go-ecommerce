#!/bin/bash
# Redis Cluster Test Script
# Tạo các thư mục cần thiết nếu chưa tồn tại
# Kiểm tra kết nối đến tất cả các node Redis và Sentinel
# Xác minh trạng thái replication
# Kiểm tra xem Sentinel nhận diện đúng master không
# Thực hiện kiểm tra tính nhất quán dữ liệu
# Kích hoạt failover thủ công để kiểm tra tính năng chuyển giao
# Hiển thị logs từ các Sentinel
# Kiểm tra nội dung cấu hình Sentinel
# Cung cấp tổng kết tình trạng của cluster

# Màu sắc cho output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[0;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

echo -e "${BLUE}============================================${NC}"
echo -e "${BLUE}     REDIS CLUSTER HEALTH CHECK SCRIPT     ${NC}"
echo -e "${BLUE}============================================${NC}"

# Kiểm tra thư mục cấu trúc
echo -e "\n${YELLOW}[1] Kiểm tra cấu trúc thư mục cần thiết...${NC}"
mkdir -p redis-mater redis-replica redis-sentinel/sentinel-1.conf redis-sentinel/sentinel-2.conf redis-sentinel/sentinel-3.conf data/redis-master-data data/redis-replica-1-data data/redis-replica-2-data data/redis-replica-3-data

# Kiểm tra kết nối đến tất cả Redis instances
echo -e "\n${YELLOW}[2] Kiểm tra kết nối đến tất cả Redis instances...${NC}"

check_redis() {
    local container=$1
    local port=$2
    local status
    
    echo -ne "   Kiểm tra $container... "
    if docker exec $container redis-cli -p $port ping 2>/dev/null | grep -q "PONG"; then
        echo -e "${GREEN}OK${NC}"
        return 0
    else
        echo -e "${RED}THẤT BẠI${NC}"
        return 1
    fi
}

# Kiểm tra Redis instances
check_redis redis-master-go-ecommerce 6379
master_ok=$?

check_redis redis-replica-1-go-ecommerce 6379
replica1_ok=$?

check_redis redis-replica-2-go-ecommerce 6379
replica2_ok=$?

check_redis redis-replica-3-go-ecommerce 6379
replica3_ok=$?

# Kiểm tra Sentinel instances
echo -e "\n${YELLOW}[3] Kiểm tra kết nối đến tất cả Sentinel instances...${NC}"

check_sentinel() {
    local container=$1
    local port=$2
    local status
    
    echo -ne "   Kiểm tra $container... "
    if docker exec $container redis-cli -p $port ping 2>/dev/null | grep -q "PONG"; then
        echo -e "${GREEN}OK${NC}"
        return 0
    else
        echo -e "${RED}THẤT BẠI${NC}"
        return 1
    fi
}

check_sentinel redis-sentinel-1-go-ecommerce 26379
sentinel1_ok=$?

check_sentinel redis-sentinel-2-go-ecommerce 26379
sentinel2_ok=$?

check_sentinel redis-sentinel-3-go-ecommerce 26379
sentinel3_ok=$?

# Kiểm tra trạng thái replication
echo -e "\n${YELLOW}[4] Kiểm tra trạng thái replication...${NC}"
if [ $master_ok -eq 0 ] && [ $replica1_ok -eq 0 ] && [ $replica2_ok -eq 0 ] && [ $replica3_ok -eq 0 ]; then
    echo -e "   Kiểm tra master info:"
    docker exec redis-master-go-ecommerce redis-cli info replication | grep role
    
    echo -e "\n   Kiểm tra replica 1 info:"
    docker exec redis-replica-1-go-ecommerce redis-cli info replication | grep -E "role|master_host|master_link_status"
    
    echo -e "\n   Kiểm tra replica 2 info:"
    docker exec redis-replica-2-go-ecommerce redis-cli info replication | grep -E "role|master_host|master_link_status"
    
    echo -e "\n   Kiểm tra replica 3 info:"
    docker exec redis-replica-3-go-ecommerce redis-cli info replication | grep -E "role|master_host|master_link_status"
else
    echo -e "${RED}   Một số Redis instances không kết nối được, bỏ qua kiểm tra replication.${NC}"
fi

# Kiểm tra sentinel có nhận diện đúng master không
echo -e "\n${YELLOW}[5] Kiểm tra Sentinel có nhận diện đúng master không...${NC}"
if [ $sentinel1_ok -eq 0 ]; then
    echo -e "   Sentinel 1 master:"
    docker exec redis-sentinel-1-go-ecommerce redis-cli -p 26379 sentinel get-master-addr-by-name mymaster
    
    echo -e "\n   Sentinel 1 slaves:"
    docker exec redis-sentinel-1-go-ecommerce redis-cli -p 26379 sentinel slaves mymaster | grep -E "ip|port"
    
    echo -e "\n   Sentinel 1 sentinels:"
    docker exec redis-sentinel-1-go-ecommerce redis-cli -p 26379 sentinel sentinels mymaster | grep -E "ip|port"
else
    echo -e "${RED}   Sentinel 1 không kết nối được, bỏ qua kiểm tra.${NC}"
fi

# Kiểm tra tính nhất quán của dữ liệu
echo -e "\n${YELLOW}[6] Kiểm tra tính nhất quán của dữ liệu...${NC}"
if [ $master_ok -eq 0 ] && [ $replica1_ok -eq 0 ] && [ $replica2_ok -eq 0 ] && [ $replica3_ok -eq 0 ]; then
    # Set a test value on master
    timestamp=$(date +%s)
    echo -e "   Đặt giá trị test trên master (timestamp: $timestamp)..."
    docker exec redis-master-go-ecommerce redis-cli set test_key "test_value_$timestamp"
    sleep 2  # Cho phép thời gian để replicate
    
    # Get values from all nodes
    master_val=$(docker exec redis-master-go-ecommerce redis-cli get test_key)
    echo -e "   Master value: ${BLUE}$master_val${NC}"
    
    replica1_val=$(docker exec redis-replica-1-go-ecommerce redis-cli get test_key)
    echo -e "   Replica 1 value: ${BLUE}$replica1_val${NC}"
    
    replica2_val=$(docker exec redis-replica-2-go-ecommerce redis-cli get test_key)
    echo -e "   Replica 2 value: ${BLUE}$replica2_val${NC}"
    
    replica3_val=$(docker exec redis-replica-3-go-ecommerce redis-cli get test_key)
    echo -e "   Replica 3 value: ${BLUE}$replica3_val${NC}"
    
    # Check consistency
    if [ "$master_val" = "$replica1_val" ] && [ "$master_val" = "$replica2_val" ] && [ "$master_val" = "$replica3_val" ]; then
        echo -e "   ${GREEN}Dữ liệu nhất quán trên tất cả các node!${NC}"
    else
        echo -e "   ${RED}Dữ liệu KHÔNG nhất quán trên tất cả các node!${NC}"
    fi
else
    echo -e "${RED}   Một số Redis instances không kết nối được, bỏ qua kiểm tra tính nhất quán dữ liệu.${NC}"
fi

# Kiểm tra khả năng failover
echo -e "\n${YELLOW}[7] Thử nghiệm failover thủ công...${NC}"
if [ $sentinel1_ok -eq 0 ]; then
    echo -e "   Master trước failover:"
    docker exec redis-sentinel-1-go-ecommerce redis-cli -p 26379 sentinel get-master-addr-by-name mymaster
    
    echo -e "\n   Kích hoạt failover thủ công..."
    docker exec redis-sentinel-1-go-ecommerce redis-cli -p 26379 sentinel failover mymaster
    
    echo -e "   Đợi 15 giây cho quá trình failover..."
    sleep 15
    
    echo -e "   Master sau failover:"
    docker exec redis-sentinel-1-go-ecommerce redis-cli -p 26379 sentinel get-master-addr-by-name mymaster
    
    echo -e "\n   Kiểm tra thông tin replication sau failover:"
    # Lấy địa chỉ IP của master mới
    new_master_ip=$(docker exec redis-sentinel-1-go-ecommerce redis-cli -p 26379 sentinel get-master-addr-by-name mymaster | head -1)
    
    for container in redis-master-go-ecommerce redis-replica-1-go-ecommerce redis-replica-2-go-ecommerce redis-replica-3-go-ecommerce
    do
        container_ip=$(docker exec $container redis-cli config get bind | tail -1 | sed 's/.*\(172[^ ]*\).*/\1/')
        role=$(docker exec $container redis-cli info replication | grep role | cut -d: -f2 | tr -d '[:space:]')
        
        echo -e "   Container $container (IP: $container_ip) role: ${BLUE}$role${NC}"
        
        if [ "$role" = "master" ]; then
            echo -e "   --> ${GREEN}Đây là master mới${NC}"
        fi
    done
else
    echo -e "${RED}   Sentinel 1 không kết nối được, bỏ qua kiểm tra failover.${NC}"
fi

# Kiểm tra logs
echo -e "\n${YELLOW}[8] Kiểm tra logs Sentinel...${NC}"
echo -e "   Lấy 5 dòng log gần nhất từ mỗi Sentinel:"

echo -e "\n   Sentinel 1 logs:"
docker logs --tail 5 redis-sentinel-1-go-ecommerce

echo -e "\n   Sentinel 2 logs:"
docker logs --tail 5 redis-sentinel-2-go-ecommerce

echo -e "\n   Sentinel 3 logs:"
docker logs --tail 5 redis-sentinel-3-go-ecommerce

# Kiểm tra cấu hình file sentinel
echo -e "\n${YELLOW}[9] Kiểm tra nội dung cấu hình Sentinel...${NC}"
if [ $sentinel1_ok -eq 0 ]; then
    echo -e "   Hiển thị file cấu hình sentinel 1 hiện tại:"
    docker exec redis-sentinel-1-go-ecommerce cat /usr/local/etc/redis/sentinel.conf | grep -v "^#" | grep -v "^$"
else 
    echo -e "${RED}   Không thể kiểm tra file cấu hình Sentinel 1.${NC}"
fi

# Kiểm tra khả năng tự động phục hồi sau khi master quay lại
echo -e "\n${YELLOW}[10] Kiểm tra master gốc sau failover...${NC}"
if [ $master_ok -eq 0 ]; then
    role=$(docker exec redis-master-go-ecommerce redis-cli info replication | grep role | cut -d: -f2 | tr -d '[:space:]')
    echo -e "   Container redis-master-go-ecommerce hiện tại có role: ${BLUE}$role${NC}"
    
    if [ "$role" = "slave" ]; then
        echo -e "   --> ${GREEN}Master gốc đã được tự động chuyển thành slave!${NC}"
        echo -e "   Thông tin replication của master gốc:"
        docker exec redis-master-go-ecommerce redis-cli info replication | grep -E "master_host|master_port"
    elif [ "$role" = "master" ]; then
        echo -e "   --> ${YELLOW}Master gốc vẫn giữ vai trò master.${NC}"
    else
        echo -e "   --> ${RED}Không xác định được role của master gốc.${NC}"
    fi
else
    echo -e "${RED}   Master gốc không kết nối được, bỏ qua kiểm tra.${NC}"
fi

# Tóm tắt kết quả
echo -e "\n${BLUE}============================================${NC}"
echo -e "${BLUE}             TỔNG KẾT KIỂM TRA             ${NC}"
echo -e "${BLUE}============================================${NC}"

echo -e "Redis Master: $([ $master_ok -eq 0 ] && echo -e "${GREEN}OK${NC}" || echo -e "${RED}THẤT BẠI${NC}")"
echo -e "Redis Replica 1: $([ $replica1_ok -eq 0 ] && echo -e "${GREEN}OK${NC}" || echo -e "${RED}THẤT BẠI${NC}")"
echo -e "Redis Replica 2: $([ $replica2_ok -eq 0 ] && echo -e "${GREEN}OK${NC}" || echo -e "${RED}THẤT BẠI${NC}")"
echo -e "Redis Replica 3: $([ $replica3_ok -eq 0 ] && echo -e "${GREEN}OK${NC}" || echo -e "${RED}THẤT BẠI${NC}")"
echo -e "Redis Sentinel 1: $([ $sentinel1_ok -eq 0 ] && echo -e "${GREEN}OK${NC}" || echo -e "${RED}THẤT BẠI${NC}")"
echo -e "Redis Sentinel 2: $([ $sentinel2_ok -eq 0 ] && echo -e "${GREEN}OK${NC}" || echo -e "${RED}THẤT BẠI${NC}")"
echo -e "Redis Sentinel 3: $([ $sentinel3_ok -eq 0 ] && echo -e "${GREEN}OK${NC}" || echo -e "${RED}THẤT BẠI${NC}")"

echo -e "\n${YELLOW}Kết luận:${NC}"
if [ $master_ok -eq 0 ] && [ $replica1_ok -eq 0 ] && [ $replica2_ok -eq 0 ] && [ $replica3_ok -eq 0 ] && [ $sentinel1_ok -eq 0 ] && [ $sentinel2_ok -eq 0 ] && [ $sentinel3_ok -eq 0 ]; then
    echo -e "${GREEN}Redis Cluster hoạt động tốt!${NC}"
else
    echo -e "${RED}Redis Cluster có vấn đề!${NC}"
    echo -e "Vui lòng kiểm tra các thành phần bị đánh dấu THẤT BẠI ở trên."
fi

echo -e "\n${BLUE}Định cấu hình Sentinel:${NC}"
echo -e "- Đảm bảo tất cả các file sentinel-1.conf, sentinel-2.conf, sentinel-3.conf đều có cấu hình như sau:"
echo -e "  * sentinel monitor mymaster redis-master-go-ecommerce 6379 2"
echo -e "  * sentinel down-after-milliseconds mymaster 5000"
echo -e "  * sentinel failover-timeout mymaster 60000"
echo -e "  * sentinel parallel-syncs mymaster 1"
echo -e "  * protected-mode no"
echo -e "  * sentinel resolve-hostnames yes"
echo -e "  * sentinel announce-hostnames yes"