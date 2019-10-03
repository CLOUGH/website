import { ILinkState, initialLinkState } from './../link/link.state';
export interface IAppState {
  link: ILinkState;
}

export const initialAppState: IAppState = {
  link: initialLinkState
}

export function getIntialState(): IAppState {
  return initialAppState;
}
