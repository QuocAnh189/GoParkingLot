import { ETypeHistory, ECardType, EVehicleType } from '@constants/enum'

export interface IIOHistory {
  id: string
  type: ETypeHistory
  card_id: string
  image_url: string
  card_type: ECardType
  vehicle_type: EVehicleType
  crop_url: string
  created_at?: string
  updated_at?: string
}

export interface IListIOHistoryRequest {
  type: string
  card_type: string
  vehicle_type: string
  date: string
  page: number
  size: number
  order_by: string
  order_desc: boolean
  take_all: boolean
}

export interface ICreateIOHistoryPayload {
  type: ETypeHistory
  rfid: string
  image: any
}

export interface IExitResponse {
  data_in: IIOHistory
  data_out: IIOHistory
  data_card: ICard
}

interface ICard {
  id: string
  owner_name: string
  card_type: string
  vehicle_type: string
  license_plate: string
  expired_date: string
}
