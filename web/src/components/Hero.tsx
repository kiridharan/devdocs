import styles from './Hero.module.css';

export default function Hero() {
    return (
        <section className={styles.hero}>
            <div className={styles.content}>
                <div className={styles.badge}>
                    <span>New: AI-Powered Generation 2.0</span>
                </div>
                <h1 className={styles.title}>
                    Generate Production-Ready<br />
                    <span className={styles.gradientText}>Documentation. In Seconds.</span>
                </h1>
                <p className={styles.subtitle}>
                    Stop writing docs. Start generating them. Instantly turn your code into comprehensive, clean, and maintainable documentation.
                </p>
                <div className={styles.actions}>
                    <button className={styles.primaryBtn}>Start Free Trial</button>
                    <button className={styles.secondaryBtn}>Watch Demo</button>
                </div>
                <div className={styles.trust}>
                    <p>Trusted by 5,000+ developers</p>
                    {/* Add logos here later */}
                </div>
            </div>
        </section>
    );
}
