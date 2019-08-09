import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { HomeRoutingModule } from './home-routing.module';
import { HomePageComponent } from './components/home-page/home-page.component';
import { CoreModule } from 'src/app/core/core.module';
import { PostPreviewComponent } from './components/post-preview/post-preview.component';
import { MomentModule } from 'ngx-moment';

@NgModule({
  declarations: [HomePageComponent, PostPreviewComponent],
  imports: [
    CommonModule,
    HomeRoutingModule,
    CoreModule,
    MomentModule
  ]
})
export class HomeModule { }
