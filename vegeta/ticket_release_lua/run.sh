#! /bin/bash
vegeta attack -rate=1000 -max-workers=100 -duration=8s --targets ./attack.txt | vegeta report