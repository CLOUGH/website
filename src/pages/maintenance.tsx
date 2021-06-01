import Head from 'next/head'
import styles from '../assets/styles/Maintenance.module.css';
import cx from 'classnames';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faEnvelope } from '@fortawesome/free-regular-svg-icons'
import { faTwitter, faGithub, faLinkedin } from '@fortawesome/free-brands-svg-icons';

export default function Maintenance() {

    return (
        <main className={styles.main}>
            <div className="container text-center">
                <h1 className="display-5">The Site Is Currently Under Construction</h1>
                <p className="lead">If you need to get in touch, please feel free to reach out using the following social media links</p>
                <div className={styles.contactIcons}>
                    <a href="https://twitter.com/warren_clough" target="_blank" >
                        <FontAwesomeIcon icon={faTwitter} />
                    </a>
                    <a href="https://github.com/CLOUGH" target="_blank" >
                        <FontAwesomeIcon icon={faGithub} />
                    </a> 
                    <a href="https://www.linkedin.com/in/warren-clough-848554a6/" target="_blank" >
                        <FontAwesomeIcon icon={faLinkedin} />
                    </a>
                    <a href="mailto:clough.warren@gmail.com" target="_blank" >
                        <FontAwesomeIcon icon={faEnvelope} />
                    </a>
                </div>
            </div>
        </main>
    )
}