# ------------------------------
# ECS cluster
# ------------------------------
resource "aws_ecs_cluster" "ecs_cluster" {
  name = "${var.project}-${var.environment}-ecs-cluster"
}

# ------------------------------
# ECS Service
# ------------------------------
resource "aws_iam_role" "ecs_task_execution_role" {
  name               = "${var.project}-${var.environment}-ecs-task-execution"
  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Principal": {
        "Service": "ecs-tasks.amazonaws.com"
      },
      "Effect": "Allow",
      "Sid": ""
    }
  ]
}
EOF
}

resource "aws_iam_role_policy_attachment" "ecs_task_execution_role_attach" {
  role       = aws_iam_role.ecs_task_execution_role.name
  policy_arn = aws_iam_policy.ecs_task_execution.arn
}

resource "aws_ecs_task_definition" "ecs_task_definition" {
  family = "${var.project}-${var.environment}-task"

  cpu                      = 256
  memory                   = 2048
  network_mode             = "awsvpc"
  requires_compatibilities = ["FARGATE"]

  container_definitions = data.template_file.container_definitions.rendered
  execution_role_arn    = aws_iam_role.ecs_task_execution_role.arn
}

data "template_file" "container_definitions" {
  template = file("./container_definitions.json")

  vars = {
    account_id  = local.account_id
    db_host     = aws_db_instance.mysql_standalone.address
    db_user     = var.db_user
    db_password = random_string.db_password.result
    db_name     = var.db_name

    ssm_agent_code = aws_ssm_activation.ssm_activation.activation_code
    ssm_agent_id   = aws_ssm_activation.ssm_activation.id
  }
}

resource "aws_cloudwatch_log_group" "cloudwatch_log_group" {
  count = length(var.app_names)
  name  = "/ecs/go-nuxt/${var.app_names[count.index]}"
}

resource "aws_ecs_service" "ecs_service" {
  name            = "${var.project}-${var.environment}-service"
  launch_type     = "FARGATE"
  desired_count   = 1
  cluster         = aws_ecs_cluster.ecs_cluster.name
  task_definition = aws_ecs_task_definition.ecs_task_definition.arn

  network_configuration {
    security_groups = [aws_security_group.alb_sg.id]
    subnets = [
      aws_subnet.private_subnet_1a.id,
      aws_subnet.private_subnet_1c.id,
      aws_subnet.private_subnet_1d.id
    ]
    assign_public_ip = true
  }

  load_balancer {
    target_group_arn = aws_lb_target_group.lb_target_group.arn
    container_name   = "nginx"
    container_port   = 80
  }
}