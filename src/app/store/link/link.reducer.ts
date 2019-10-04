import { LinkActions, ELinkActions } from './link.action';
import { initialLinkState, ILinkState } from './link.state';


export function linkReducers(
  state = initialLinkState,
  action: LinkActions
): ILinkState {
  switch (action.type) {
    case ELinkActions.GetLinksSuccess: {
      return {
        ...state,
        links: action.payload
      };
    }

    default:
      return state;
  }
}
