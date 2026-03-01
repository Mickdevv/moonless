export type JwtPayload = {
  iss: string
  sub: string
  exp: number
  iat: number
  role: string
  email: string
}
