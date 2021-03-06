// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

data "oci_identity_availability_domain" "ad" {
  compartment_id = "${var.tenancy_ocid}"
  ad_number      = 1
}

data "oci_core_instance_configuration" TFInstanceConfigurationDatasource {
  instance_configuration_id = "${oci_core_instance_configuration.TFInstanceConfiguration.id}"
}

data "oci_core_instance_configurations" TFInstanceConfigurationDatasources {
  compartment_id = "${var.compartment_ocid}"

  filter {
    name   = "id"
    values = ["${oci_core_instance_configuration.TFInstanceConfiguration.id}"]
  }
}

data "oci_core_instance_pool" "TFInstancePoolDatasource" {
  instance_pool_id = "${oci_core_instance_pool.TFInstancePool.id}"
}

data "oci_core_instance_pools" "TFInstancePoolDatasources" {
  compartment_id = "${var.compartment_ocid}"
  display_name   = "TFInstancePool"
  state          = "RUNNING"

  filter {
    name   = "id"
    values = ["${oci_core_instance_pool.TFInstancePool.id}"]
  }
}

data "oci_core_instance_pool_instances" "TFInstancePoolInstanceDatasources" {
  compartment_id   = "${var.compartment_ocid}"
  instance_pool_id = "${oci_core_instance_pool.TFInstancePool.id}"
  display_name     = "TFInstancePool"
}
