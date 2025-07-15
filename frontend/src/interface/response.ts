export interface BaseResponse<T> {
  code: number
  message: string
  error?: string
  data: T
  pagination?: Pagination
}

export interface Pagination {
  page: number
  size: number
  totalPage: number
  totalData: number
}

export interface GetVoucherListResponse {
  id: number
  crewId: string
  crewName: string
  flightNumber: string
  flightDate: string
  aircraftType: string
  seat: string[]
  createdAt: string
}

export interface CheckVoucherResponse {
  exists: boolean
}

export interface GenerateVoucherResponse {
  success: boolean
  seats: string[]
}
