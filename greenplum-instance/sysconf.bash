#!/bin/bash

#install required deb
sub_path=/greenplum-for-kubernetes/greenplum-instance
for deb in `cat ${sub_path}/greenplum_instance_apt_install_order.txt`;do
  deb=${deb#*_}
  dpkg -i ${sub_path}/greenplum-instance-debs/${deb}
done

