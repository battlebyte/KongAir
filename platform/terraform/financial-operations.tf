resource "konnect_gateway_control_plane" "financial_ops_control_plane" {
  name         = "Financial Operations Control Plane"
  description  = "financial Operations Business Domain"
  cluster_type = "CLUSTER_TYPE_CONTROL_PLANE"
  auth_type    = "pinned_client_certs"
  cloud_gateway = false
}

output "financial_ops_control_plane_id" {
  value = konnect_gateway_control_plane.financial_ops_control_plane.id
}

resource "konnect_team" "financial_ops_team" {
  name        = "Financial Operations Team"
  description = "Financial Operations Team managed by TF"

  labels = {
    example = "here"
  }
}

resource "konnect_team_role" "financial_ops_team_cp_role" {
  entity_id        = konnect_gateway_control_plane.financial_ops_control_plane.id
  entity_region    = "eu"
  entity_type_name = "Control Planes"
  role_name        = "Admin"
  team_id          = konnect_team.financial_ops_team.id
}