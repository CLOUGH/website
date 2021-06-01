import { PageState } from './page.state';
import { IAppState } from '../app/app.state';
import { createSelector } from '@ngrx/store';

export const selectPageState = (state: IAppState) => state.pages;

export const selectPages = createSelector(
  selectPageState,
  (state: PageState) => {
    return state.pages;
});
