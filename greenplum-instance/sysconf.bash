#!/bin/bash

#install required deb
sub_path=/tools
for deb in `cat ${sub_path}/greenplum_instance_apt_install_order.txt`;do
  deb=${deb#*_}
  dpkg -i ${sub_path}/${deb}
done

