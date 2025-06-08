
import type { Phone } from "@/api/phone";

export function processPhoneInfo(PhoneList: Phone[] ): Map<string,string>[] {
    const processedInfo: Map<string, string>[] = [];
    for (const phone of PhoneList) {
        const phoneInfo = new Map<string, string>();
        phoneInfo.set('serial', phone.serial);
        phoneInfo.set('manufacturer', phone.manufacturer);
        phoneInfo.set('model', phone.model);
        phoneInfo.set('androidVersion', phone.androidVersion);
        phoneInfo.set('cpuabi', phone.cpuabi);
        phoneInfo.set('marketName', phone.marketName);
        phoneInfo.set('marketNameSymbol', phone.marketNameSymbol);

        processedInfo.push(phoneInfo);
    }
  return processedInfo;
}