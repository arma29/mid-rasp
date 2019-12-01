#!/bin/bash

localhost='127.0.0.1'
r_client='SensorR'

for i in {1..31}
do
    ($r_client guest guest $localhost)
done

echo 'DONE'