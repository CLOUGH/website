import { ILink } from './link';

export interface ILinkState {
  links: ILink[]
}


export const initialLinkState: ILinkState = {
  links: []
}
