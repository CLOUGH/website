import { Action } from '@ngrx/store';
import { ILink } from './link';

export enum ELinkActions {
    GetLinks = '[Link] Get Links',
    GetLinksSuccess = '[Link] Get Links Success'
}

export class GetLinks implements Action {
    public readonly type = ELinkActions.GetLinks;
}

export class GetLinksSuccess implements Action {
    public readonly type = ELinkActions.GetLinksSuccess;
    constructor(public payload: ILink[]){}
}


export type LinkActions = GetLinks | GetLinksSuccess;
