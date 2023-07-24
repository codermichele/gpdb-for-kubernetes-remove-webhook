#!/bin/bash

#install required deb
sub_path=/greenplum-for-kubernetes/greenplum-operator
for deb in `cat ${sub_path}/greenplum_operator_apt_install_order.txt`;do
  deb=${deb#*_}
  dpkg -i ${sub_path}/${deb}
done
