export interface baseResponse {
    message: string;
    status: boolean;
}

export interface nodesResponse extends baseResponse {
    nodes: Map<string, string>[];
}