import { PageState, initialPagesState } from '../page/page.state';
export interface IAppState {
  pages: PageState;
}

export const initialAppState: IAppState = {
  pages: initialPagesState
}

export function getIntialState(): IAppState {
  return initialAppState;
}
