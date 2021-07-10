import React from "react";
import cx from 'classnames';
import { useRouter } from "next/dist/client/router";
import styled from 'styled-components';
import Link from 'next/link';


interface Props {}


const StyledNavbar = styled.div`
    font-size: 12px;
`;

export const MainNavbar: React.FC<Props> = (props) => {
    const router = useRouter();

    return (
        <StyledNavbar >
            <nav className="navbar navbar-expand-lg navbar-light">

                <div className="container">
                    <a className="navbar-brand" href="#">WC</a>
                    <button className="navbar-toggler" 
                        type="button" 
                        data-bs-toggle="collapse" 
                        data-bs-target="#navbarSupportedContent" 
                        aria-controls="navbarSupportedContent" 
                        aria-expanded="false" 
                        aria-label="Toggle navigation">
                        <span className="navbar-toggler-icon"></span>
                    </button>
                    <div className="collapse navbar-collapse" id="navbarSupportedContent">
                        <ul className="navbar-nav me-auto mb-2 mb-lg-0"></ul>
                        <ul className="navbar-nav mb-2 mb-lg-0">
                            <li className="nav-item">
                                <Link href="/">
                                    <a className="nav-link" >
                                        About
                                    </a>
                                </Link>
                            </li>
                            <li className="nav-item">
                                <Link href="/portfolio">
                                    <a className="nav-link" >Portfolio</a>
                                </Link>
                            </li>
                            <li className="nav-item">
                                <Link href="/contact">
                                    <a className="nav-link">Contact</a>
                                </Link>
                            </li>
                        </ul>
                    </div>
                </div>
            </nav>
        </StyledNavbar>
    )
};