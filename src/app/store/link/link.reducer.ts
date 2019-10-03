import { LinkActions, ELinkActions } from './link.action';
import { initialLinkState, ILinkState } from './link.state';


export const linkReducers = (
  state = initialLinkState,
  action: LinkActions
): ILinkState => {
  console.log({type: action.type});
  switch (action.type) {
    case ELinkActions.GetLinksSuccess: {
      console.log({
        ...state,
        links: action.payload
      },'test');
      return {
        ...state,
        links: action.payload
      };
    }

    default:
      return state;
  }
}
