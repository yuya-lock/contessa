# ------------------------------
# Certificate
# ------------------------------
# for tokyo region
resource "aws_acm_certificate" "tokyo_cert" {
  domain_name       = "*.${var.domain}"
  validation_method = "DNS"

  tags = {
    Name    = "${var.project}-${var.environment}-wildcard-sslcert"
    Project = var.project
    Env     = var.environment
  }

  lifecycle {
    create_before_destroy = true
  }
}

resource "aws_route53_record" "route53_acm_dns_resolve" {
  depends_on = [aws_acm_certificate.tokyo_cert]

  for_each = {
    for dvo in aws_acm_certificate.tokyo_cert.domain_validation_options : dvo.domain_name => {
      name   = dvo.resource_record_name
      type   = dvo.resource_record_type
      record = dvo.resource_record_value
    }
  }

  allow_overwrite = true
  zone_id         = aws_route53_zone.route53_zone.id
  name            = each.value.name
  type            = each.value.type
  ttl             = 600
  records         = [each.value.record]
}

resource "aws_acm_certificate_validation" "tokyo_cert_valid" {
  certificate_arn         = aws_acm_certificate.tokyo_cert.arn
  validation_record_fqdns = [for record in aws_route53_record.route53_acm_dns_resolve : record.fqdn]
}

# for virginia region
resource "aws_acm_certificate" "virginia_cert" {
  provider = aws.virginia

  domain_name       = var.domain
  validation_method = "DNS"

  tags = {
    Name    = "${var.project}-${var.environment}-wildcard-sslcert"
    Project = var.project
    Env     = var.environment
  }

  lifecycle {
    create_before_destroy = true
  }
}

# resource "aws_route53_record" "route53_acm_dns_virginia_resolve" {
#   depends_on = [aws_acm_certificate.virginia_cert]

#   for_each = {
#     for dvo in aws_acm_certificate.virginia_cert.domain_validation_options : dvo.domain_name => {
#       name   = dvo.resource_record_name
#       type   = dvo.resource_record_type
#       record = dvo.resource_record_value
#     }
#   }

#   zone_id = aws_route53_zone.route53_zone.id
#   name    = each.value.name
#   type    = each.value.type
#   ttl     = 600
#   records = [each.value.record]
# }

# resource "aws_acm_certificate_validation" "virginia_cert_valid" {
#   provider = aws.virginia

#   certificate_arn         = aws_acm_certificate.virginia_cert.arn
#   validation_record_fqdns = [for record in aws_route53_record.route53_acm_dns_virginia_resolve : record.fqdn]
# }