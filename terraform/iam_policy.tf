# ------------------------------
# IAM Policy
# ------------------------------
# for SSM activation
# resource "aws_iam_policy" "update_ssm_service_policy" {
#   name        = "${var.project}-${var.environment}-update-ssm-service-policy"
#   description = "Update ssm service policy"
#   policy      = data.aws_iam_policy_document.update_ssm_service_policy.json
# }

# data "aws_iam_policy_document" "update_ssm_service_policy" {
#   statement {
#     effect = "Allow"
#     actions = [
#       "ssm:GetServiceSetting"
#     ]
#     resources = ["*"]
#   }

#   statement {
#     effect = "Allow"
#     actions = [
#       "ssm:ResetServiceSetting",
#       "ssm:UpdateServiceSetting"
#     ]
#     resources = ["arn:aws:ssm:ap-northeast-1:${data.aws_caller_identity.user.account_id}:servicesetting/ssm/managed-instance/activation-tier"]
#   }
# }

# data "aws_caller_identity" "user" {}

# locals {
#   account_id = data.aws_caller_identity.user.account_id
# }

# # for ECS
# resource "aws_iam_policy" "ecs_task_execution" {
#   name   = "${var.project}-${var.environment}-ecs-task-exection"
#   policy = data.aws_iam_policy_document.ecs_task_exection.json
# }

# data "aws_iam_policy" "ecs_task_exection_role_policy" {
#   arn = "arn:aws:iam::aws:policy/service-role/AmazonECSTaskExecutionRolePolicy"
# }

# data "aws_iam_policy_document" "ecs_task_exection" {
#   source_json = data.aws_iam_policy.ecs_task_exection_role_policy.policy

#   statement {
#     effect    = "Allow"
#     actions   = ["ssm:GetParameters", "kms:Decrypt"]
#     resources = ["*"]
#   }
# }