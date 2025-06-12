import request from "@/utils/request";

export interface Phone {
  id: string;
  serial: string; 
  manufacturer: string;
  model: string;
  androidVersion: string;
  cpuabi: string;
  marketName: string;
  marketNameSymbol: string;
}

export const addPhone = (phone: Phone) => {
    return request({
        url: '/api/v1/phone/add',
        method: 'post',
        data: phone
    });
}

export const getAllPhoneInfo = () => {
    return request({
        url: '/api/v1/phone/find',
        method: 'get'
    });
}

export const updataPhone = (phone: Phone) => {
    return request({
        url: '/api/v1/phone/update',
        method: 'put',
        data: phone
    });
}

export const deletePhone = (serial: string) => {    
    return request({
        url: '/api/v1/phone/remove'+`?serial=${serial}`,
        method: 'delete',
    });
}