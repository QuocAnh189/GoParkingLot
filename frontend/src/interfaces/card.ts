import { ECardType, EVehicleType } from "@constants/enum";

export interface ICard {
  id: string;
  rfid: string;
  owner_name: string;
  card_type: ECardType;
  vehicle_type: EVehicleType;
  license_plate: string;
  expired_date: string;
  last_io_history_id: string;
  created_at?: string;
  updated_at?: string;
}

export interface IListCardRequest {
  search: string;
  card_type: string;
  vehicle_type: string;
  page: number;
  size: number;
  order_by: string;
  order_desc: boolean;
  take_all: boolean;
}

export interface CreateCardPayload {
  rfid: string;
  owner_name: string;
  card_type: ECardType;
  vehicle_type: EVehicleType;
  license_plate: string;
  expired_date: string;
}

export interface UpdateCardPayload {
  id: string;
  rfid: string;
  owner_name: string;
  card_type: ECardType;
  vehicle_type: EVehicleType;
  license_plate: string;
  expired_date: string;
  last_io_history_id: null | string;
}
