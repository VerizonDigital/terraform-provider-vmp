# Copyright Verizon Media, Licensed under the terms of the Apache 2.0 license . See LICENSE file in project root for terms.
terraform {
  required_providers {
    vmp = {
      version = "0.0.8"
      source = "github.com/terraform-providers/vmp"
    }
  }
}
##########################################
# Variables
##########################################
variable "credentials" {
  type = object ({
    api_token = string
    ids_client_secret = string
    ids_client_id = string
    ids_scope = string
    api_address = string
    ids_address = string
  })
}

variable "new_admin_user" {
  type = object({
    customer_account_number = string
    first_name = string
    last_name = string
    email = string
    is_admin = bool
  })
  default = {
    customer_account_number = ""
    first_name = "admin"
    last_name = "user"
    email = ""
    is_admin = true
  }
}

##########################################
# Providers
##########################################
provider "vmp" {
    api_token = var.credentials.api_token
    ids_client_secret = var.credentials.ids_client_secret
    ids_client_id = var.credentials.ids_client_id
    ids_scope = var.credentials.ids_scope
    api_address = var.credentials.api_address
    ids_address = var.credentials.ids_address
    partner_id = 3249
    partner_user_id = 7613
}
##########################################
# Resources
##########################################
resource "vmp_customer_user" "test_customer_noadmin1" {
  account_number = "D9127"
  first_name = var.new_admin_user.first_name
  last_name = var.new_admin_user.last_name
  email = "admin+1@test20252021-3.com"
  is_admin = false
}
