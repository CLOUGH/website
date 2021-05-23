import { Action } from '@ngrx/store';
import { Page } from '../../core/models/page';

export enum EPageActions {
    GetPages = '[Page] Get Pages',
    GetPagesSuccess = '[Page] Get Pages Success'
}

export class GetPages implements Action {
    public readonly type = EPageActions.GetPages;
}

export class GetPagesSuccess implements Action {
    public readonly type = EPageActions.GetPagesSuccess;
    constructor(public payload: Page[]){}
}


export type PageActions = GetPages | GetPagesSuccess;
