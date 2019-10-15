import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { ApolloBoostModule, ApolloBoost  } from 'apollo-angular-boost';


@NgModule({
  declarations: [],
  imports: [
    CommonModule,
    ApolloBoostModule
  ]
})
export class GraphqlModule {

  constructor(apolloBoost: ApolloBoost) {
    apolloBoost.create({
      uri: 'http://localhost:4000/graphql'
    });
  }

}
