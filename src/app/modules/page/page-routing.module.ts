import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { PageDetailComponent } from './page-detail/page-detail.component';
import { PageResolverService } from './page-resolver.service';


const routes: Routes = [
  {path: '**', component: PageDetailComponent , resolve: { page: PageResolverService}}
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class PageRoutingModule { }
