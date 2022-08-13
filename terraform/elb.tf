# ------------------------------
# ALB
# ------------------------------
resource "aws_lb" "alb" {
  name               = "${var.project}-${var.environment}-app-alb"
  internal           = false
  load_balancer_type = "application"
  security_groups = [
    aws_security_group.alb_sg.id
  ]
  subnets = [
    aws_subnet.public_subnet_1a.id,
    aws_subnet.public_subnet_1c.id,
    aws_subnet.public_subnet_1d.id,
  ]
}

resource "aws_lb_listener" "alb_listener_http" {
  load_balancer_arn = aws_lb.alb.arn
  port              = 80
  protocol          = "HTTP"

  default_action {
    type = "redirect"

    redirect {
      port        = "443"
      protocol    = "HTTPS"
      status_code = "HTTP_301"
    }
  }
}

resource "aws_lb_listener" "alb_listener_https" {
  load_balancer_arn = aws_lb.alb.arn
  port              = 443
  protocol          = "HTTPS"
  certificate_arn   = aws_acm_certificate.tokyo_cert.arn

  default_action {
    type = "fixed-response"

    fixed_response {
      content_type = "text/plain"
      status_code  = "200"
      message_body = "ok"
    }
  }
}

resource "aws_lb_listener_rule" "http_rule" {
  listener_arn = aws_lb_listener.alb_listener_http.arn

  action {
    type             = "forward"
    target_group_arn = aws_lb_target_group.lb_target_group.id
  }

  condition {
    path_pattern {
      values = ["*"]
    }
  }
}

resource "aws_lb_listener_rule" "https_rule" {
  listener_arn = aws_lb_listener.alb_listener_https.arn

  action {
    type             = "forward"
    target_group_arn = aws_lb_target_group.lb_target_group.id
  }

  condition {
    path_pattern {
      values = ["*"]
    }
  }
}

# ------------------------------
# target group
# ------------------------------
resource "aws_lb_target_group" "lb_target_group" {
  name = "${var.project}-${var.environment}-tg"

  vpc_id      = aws_vpc.vpc.id
  port        = 80
  protocol    = "HTTP"
  target_type = "ip"

  health_check {
    port = 80
    path = "/health"
  }
}