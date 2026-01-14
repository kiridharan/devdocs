import Link from 'next/link';
import styles from './Navbar.module.css';

export default function Navbar() {
    return (
        <nav className={styles.nav}>
            <div className={styles.container}>
                <div className={styles.logo}>
                    <Link href="/">DevDocs</Link>
                </div>
                <div className={styles.links}>
                    <Link href="#features">Features</Link>
                    <Link href="#how-it-works">How it Works</Link>
                    <Link href="#pricing">Pricing</Link>
                </div>
                <div className={styles.cta}>
                    <Link href="/login" className={styles.loginBtn}>Login</Link>
                    <Link href="/signup" className={styles.signupBtn}>Get Started</Link>
                </div>
            </div>
        </nav>
    );
}
