export interface CheckVoucherRequest {
  flightNumber: string
  date: string
}

export interface GenerateVoucherRequest {
  id: string
  name: string
  flightNumber: string
  date: string
  aircraft: string
}