import { GetServerSideProps } from 'next'
import Head from 'next/head'
import Image from 'next/image';
import styled from 'styled-components';

import { MainNavbar } from '../components/main-navbar/main-navbar'

const StyledHeroSection = styled.div`

  .description {
    display: flex;
    flex-direction: column;
    justify-content: center;
  }
  .call-to-action {
    text-align: center;
    button {
      margin-left: 15px;
    }
  }
`;
export default function Home() {
  return (
    <div>
      <Head>
        <title>Warren Clough</title>
        <meta name="description" content="Generated by create next app" />
        <link rel="icon" href="/favicon.ico" />
      </Head>

      <MainNavbar />

      <main className="">
        <StyledHeroSection className="container mb-5 mt-5" >
           <div className="row justify-content-between">
            <div className="profile-image text-md-end text-center  text-sm-center col-12  col-md-4 order-md-last">
              <Image src="/images/DSC00989_edited.jpg"  width="280" height="350px" objectFit="cover"/>
            </div>
            <div className="description col-md-7 col-12">
              <p>Hi,  my name is Warren Clough</p>
              <h3>I'm a Software Engineer with a background in Financial Systems and Human Resource Management tools. </h3>
              <div className="call-to-action">
                <button className="btn btn-default btn-sm btn-outline-secondary">Let Chat</button>
                <button className="btn btn-sm btn-secondary">About</button>
              </div>
            </div>
          </div>
        </StyledHeroSection>

        {/* <StyledSection className="bg-light pt-5 pt-5">
          <div className="container">
            <h4>A Bit About Me</h4>
            <div className="row">
              <div className="col-12 col-md-6">
                <p>Lorem ipsum dolor sit amet consectetur adipisicing elit. Voluptatem, totam. Autem hic dolorem, quaerat rem error deserunt voluptatum necessitatibus odit mollitia corporis aliquam velit eaque rerum explicabo, reprehenderit vero natus.</p>
                <p>Lorem ipsum dolor sit amet consectetur adipisicing elit. Voluptatem, totam. Autem hic dolorem, quaerat rem error deserunt voluptatum necessitatibus odit mollitia corporis aliquam velit eaque rerum explicabo, reprehenderit vero natus.</p>
              </div>
              <div className="col-12 col-md-6"></div>
            </div>

          </div>
        </StyledSection>
        <StyledSection className=" pt-5 pt-5">
          <div className="container">
            <h4>What I Do</h4>
            <div className="row">
              <div className="col-12 col-md-6">
                <p>Lorem ipsum dolor sit amet consectetur adipisicing elit. Voluptatem, totam. Autem hic dolorem, quaerat rem error deserunt voluptatum necessitatibus odit mollitia corporis aliquam velit eaque rerum explicabo, reprehenderit vero natus.</p>
                <p>Lorem ipsum dolor sit amet consectetur adipisicing elit. Voluptatem, totam. Autem hic dolorem, quaerat rem error deserunt voluptatum necessitatibus odit mollitia corporis aliquam velit eaque rerum explicabo, reprehenderit vero natus.</p>
              </div>
              <div className="col-12 col-md-6"></div>
            </div>

          </div>
        </StyledSection> */}
      </main>
    </div>
  )
}
