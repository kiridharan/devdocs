import Link from 'next/link';
import styles from './Footer.module.css';

export default function Footer() {
    return (
        <footer className={styles.footer}>
            <div className={styles.ctaSection}>
                <h2>Ready to automate your documentation?</h2>
                <div className={styles.actions}>
                    <button className={styles.primaryBtn}>Start Generating Docs Now</button>
                    <p className={styles.note}>No credit card required</p>
                </div>
            </div>

            <div className={styles.linksSection}>
                <div className={styles.logo}>DevDocs</div>
                <div className={styles.copyright}>
                    &copy; {new Date().getFullYear()} DevDocs Inc. All rights reserved.
                </div>
                <div className={styles.socials}>
                    <Link href="#">Twitter</Link>
                    <Link href="#">GitHub</Link>
                    <Link href="#">Discord</Link>
                </div>
            </div>
        </footer>
    );
}
