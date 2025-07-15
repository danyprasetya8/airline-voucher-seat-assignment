import type { BaseResponse, GetVoucherListResponse, CheckVoucherResponse, GenerateVoucherResponse } from '@/interface/response';
import type { CheckVoucherRequest, GenerateVoucherRequest } from '@/interface/request'
import axios, { type AxiosResponse } from 'axios';

export const getVoucherList = async (page: number, size: number): Promise<AxiosResponse<BaseResponse<GetVoucherListResponse[]>>> => {
  const qs: URLSearchParams = new URLSearchParams();
  qs.append('page', page+'');
  qs.append('size', size+'');
  return await axios.get(`/api/v1/voucher?${qs.toString()}`)
}

export const checkVoucher = async (req: CheckVoucherRequest): Promise<AxiosResponse<BaseResponse<CheckVoucherResponse>>> => {
  return axios.post('/api/v1/voucher/check', req)
}

export const generateVoucher = async (req: GenerateVoucherRequest): Promise<AxiosResponse<BaseResponse<GenerateVoucherResponse>>> => {
  return axios.post('/api/v1/voucher/generate', req)
}
