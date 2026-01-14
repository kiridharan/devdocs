import styles from './SocialProof.module.css';

export default function SocialProof() {
    const testimonials = [
        {
            quote: "DevDocs saved our team 20 hours a week. It's magic.",
            author: "Alex Chen",
            role: "Lead Engineer @ TechFlow"
        },
        {
            quote: "Finally, a doc tool that actually understands code.",
            author: "Sarah Jones",
            role: "CTO @ StartupX"
        },
        {
            quote: "The best developer experience upgrade we've made this year.",
            author: "Michael Brown",
            role: "Senior Dev @ EnterpriseCo"
        }
    ];

    return (
        <section className={styles.section}>
            <div className={styles.container}>
                <div className={styles.header}>
                    <h2>Loved by Developers</h2>
                </div>
                <div className={styles.grid}>
                    {testimonials.map((t, i) => (
                        <div key={i} className={styles.card}>
                            <p className={styles.quote}>"{t.quote}"</p>
                            <div className={styles.author}>
                                <div className={styles.avatar} /> {/* Placeholder */}
                                <div>
                                    <div className={styles.name}>{t.author}</div>
                                    <div className={styles.role}>{t.role}</div>
                                </div>
                            </div>
                        </div>
                    ))}
                </div>
            </div>
        </section>
    );
}
