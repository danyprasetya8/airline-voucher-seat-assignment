import type { BaseResponse, GetVoucherListResponse, CheckVoucherResponse, GenerateVoucherResponse } from '@/interface/response';
import type { CheckVoucherRequest, GenerateVoucherRequest } from '@/interface/request'

export const getVoucherList = async (page: number, size: number): Promise<BaseResponse<GetVoucherListResponse[]>> => {
  const qs: URLSearchParams = new URLSearchParams();
  qs.append('page', page+'');
  qs.append('size', size+'');
  const f = await fetch(`/api/v1/voucher?${qs.toString()}`)
  return f.json()
}

export const checkVoucher = async (req: CheckVoucherRequest): Promise<BaseResponse<CheckVoucherResponse>> => {
  const f = await fetch('/api/v1/voucher/check', {
    method: 'POST',
    body: JSON.stringify(req)
  })
  return f.json()
}

export const generateVoucher = async (req: GenerateVoucherRequest): Promise<BaseResponse<GenerateVoucherResponse>> => {
  const f = await fetch('/api/v1/voucher/generate', {
    method: 'POST',
    body: JSON.stringify(req)
  })
  return f.json()
}
