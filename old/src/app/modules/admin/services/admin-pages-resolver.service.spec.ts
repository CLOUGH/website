import { TestBed } from '@angular/core/testing';

import { AdminPagesResolverService } from './admin-pages-resolver.service';

describe('AdminPagesResolverService', () => {
  beforeEach(() => TestBed.configureTestingModule({}));

  it('should be created', () => {
    const service: AdminPagesResolverService = TestBed.get(AdminPagesResolverService);
    expect(service).toBeTruthy();
  });
});
