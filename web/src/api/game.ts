import request from "@/utils/request";


export interface Game {
  id: String;
  packahe: String;
  type: String;
  case_type: String[];
  game_engine: String;
  game_url: String;
  game_name: String;
  game_type: String;
  game_id: number;
  status: boolean;
}

export interface SearchGame {
  game_type: string;
  case_type: string;
  game_name: string;
}

export const findAllWxGame = () => {
  return request({
    url: '/api/game/find_wx',
    method: 'get'
  });
}

export const searchGame = (data: SearchGame) => {
  return request({
    url: '/api/game/find_by_type',
    method: 'post',
    data: data
});
}
