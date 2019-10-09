import { ILinkState } from './link.state';
import { IAppState } from '../app/app.state';
import { createSelector } from '@ngrx/store';

export const selectLinkState = (state: IAppState) => state.link;

export const selectLinks = createSelector(
  selectLinkState,
  (state: ILinkState) => {
    return state.links;
});
