# ------------------------------
# SSM Activation
# ------------------------------
resource "aws_ssm_activation" "ssm_activation" {
  name               = "${var.project}-${var.environment}-ssm-activation"
  iam_role           = aws_iam_role.ssm_ec2_run_role.id
  registration_limit = "5"
  depends_on = [
    aws_iam_role_policy_attachment.ssm_ec2_run_role_attach_1,
    aws_iam_role_policy_attachment.ssm_ec2_run_role_attach_2,
    aws_iam_role_policy_attachment.ssm_ec2_run_role_attach_3,
    aws_iam_role_policy_attachment.ssm_ec2_run_role_attach_4,
  ]
}

resource "aws_iam_role" "ssm_ec2_run_role" {
  name               = "${var.project}-${var.environment}-ssm-ec2-run-role"
  assume_role_policy = data.aws_iam_policy_document.ssm_ec2_run_role.json
}

data "aws_iam_policy_document" "ssm_ec2_run_role" {
  statement {
    effect = "Allow"
    actions = [
      "sts:AssumeRole"
    ]
    principals {
      type        = "Service"
      identifiers = ["ssm.amazonaws.com"]
    }
  }
}

resource "aws_iam_role_policy_attachment" "ssm_ec2_run_role_attach_1" {
  role       = aws_iam_role.ssm_ec2_run_role.name
  policy_arn = aws_iam_policy.update_ssm_service_policy.arn
}

resource "aws_iam_role_policy_attachment" "ssm_ec2_run_role_attach_2" {
  role       = aws_iam_role.ssm_ec2_run_role.name
  policy_arn = "arn:aws:iam::aws:policy/AmazonSSMManagedInstanceCore"
}

resource "aws_iam_role_policy_attachment" "ssm_ec2_run_role_attach_3" {
  role       = aws_iam_role.ssm_ec2_run_role.name
  policy_arn = "arn:aws:iam::aws:policy/AmazonSSMDirectoryServiceAccess"
}

resource "aws_iam_role_policy_attachment" "ssm_ec2_run_role_attach_4" {
  role       = aws_iam_role.ssm_ec2_run_role.name
  policy_arn = "arn:aws:iam::aws:policy/CloudWatchAgentServerPolicy"
}