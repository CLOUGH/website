
import { ActionReducerMap } from '@ngrx/store';
import { IAppState } from './app.state';
import { linkReducers } from '../link/link.reducer';

export const appReducers: ActionReducerMap<IAppState, any> = {
  link: linkReducers
}
