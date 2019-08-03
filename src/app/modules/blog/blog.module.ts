import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { BlogRoutingModule } from './blog-routing.module';
import { BlogListingsPageComponent } from './pages/blog-listings-page/blog-listings-page.component';
import { BlogDetailPageComponent } from './pages/blog-detail-page/blog-detail-page.component';


@NgModule({
  declarations: [BlogListingsPageComponent, BlogDetailPageComponent],
  imports: [
    CoreModule,
    CommonModule,
    BlogRoutingModule
  ]
})
export class BlogModule { }
