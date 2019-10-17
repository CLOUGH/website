
import { ActionReducerMap } from '@ngrx/store';
import { IAppState } from './app.state';
import { pageReducers } from '../page/page.reducer';

export const appReducers: ActionReducerMap<IAppState, any> = {
  pages: pageReducers
}
