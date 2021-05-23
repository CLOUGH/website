import { Page } from '../../core/models/page';

export interface PageState {
  pages: Page[]
}


export const initialPagesState: PageState = {
  pages: []
}
