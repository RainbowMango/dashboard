import axios from 'axios';

const baseURL: string = "/api/v1";
export const karmadaClient = axios.create({
    baseURL,
})

export interface IResponse<Data = {}> {
    code: number;
    message: string;
    data: Data;
}

export type Labels = Record<string, string>
export type Annotations = Record<string, string>

export interface ObjectMeta {
    name: string
    namespace: string
    labels: Labels
    annotations: Annotations
    creationTimestamp: string
    uid: string
}

export interface TypeMeta {
    kind: string
    scalable: boolean
    restartable: boolean
}
export type Selector =  Record<string, string>


export interface RollingUpdateStrategy {
    maxSurge: string
    maxUnavailable: string
}