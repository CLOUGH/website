import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { PostRoutingModule } from './post-routing.module';
import { PostPreviewComponent } from './post-preview/post-preview.component';
import { SharedModule } from 'src/app/shared/shared.module';
import { MomentModule } from 'ngx-moment';


@NgModule({
  declarations: [PostPreviewComponent],
  imports: [
    CommonModule,
    PostRoutingModule,
    SharedModule,
    MomentModule,
    SharedModule

  ],
  exports: [PostPreviewComponent]
})
export class PostModule { }
