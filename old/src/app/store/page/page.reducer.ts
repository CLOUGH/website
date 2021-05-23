import { PageActions, EPageActions } from './page.action';
import { initialPagesState, PageState } from './page.state';


export function pageReducers(
  state = initialPagesState,
  action: PageActions
): PageState {
  switch (action.type) {
    case EPageActions.GetPagesSuccess: {
      return {
        ...state,
        pages: action.payload
      };
    }

    default:
      return state;
  }
}
