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
        url: '/api/phone/add',
        method: 'post',
        data: phone
    });
}

export const getAllPhoneInfo = () => {
    return request({
        url: '/api/phone/find',
        method: 'get'
    });
}

export const updataPhone = (phone: Phone) => {
    return request({
        url: '/api/phone/update',
        method: 'put',
        data: phone
    });
}

export const deletePhone = (serial: string) => {    
    return request({
        url: '/api/phone/remove'+`?serial=${serial}`,
        method: 'delete',
    });
}