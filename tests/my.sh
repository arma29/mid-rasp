#!/bin/bash

localhost='localhost'
my_sensor='SensorMy'

for i in {1..3}
do
    ($my_sensor)
    sleep 10s
done

echo 'DONE'