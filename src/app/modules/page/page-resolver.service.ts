import { PageService } from './page.service';
import { Injectable } from '@angular/core';
import { Resolve, ActivatedRouteSnapshot, RouterStateSnapshot, Router } from '@angular/router';
import { of } from 'rxjs';
import { map } from 'rxjs/operators';

@Injectable({
  providedIn: 'root'
})
export class PageResolverService implements Resolve<any> {

  constructor(private pageService: PageService, private router: Router) { }

  resolve(route: ActivatedRouteSnapshot, state: RouterStateSnapshot) {
    const path = route.url.join('/');
    return this.pageService.findPage(path).pipe(map(page => {
      if (!page) {
        this.router.navigateByUrl('error/404');
        return null;
      }
      return page;
    }));
  }
}
