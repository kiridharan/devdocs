import styles from './ValueProp.module.css';

export default function ValueProp() {
    const cards = [
        {
            title: "Lightning Fast",
            desc: "Generate full API references and guides in under 30 seconds.",
            icon: "⚡"
        },
        {
            title: "Always Up-to-Date",
            desc: "Auto-syncs with your code changes. Never worry about stale docs.",
            icon: "🔄"
        },
        {
            title: "Developer Friendly",
            desc: "CLI-first approach, integrates directly into your CI/CD pipeline.",
            icon: "💻"
        },
        {
            title: "Standardized Quality",
            desc: "Consistent format and tone across your entire project.",
            icon: "✨"
        }
    ];

    return (
        <section className={styles.section}>
            <div className={styles.container}>
                {cards.map((card, index) => (
                    <div key={index} className={styles.card}>
                        <div className={styles.icon}>{card.icon}</div>
                        <h3>{card.title}</h3>
                        <p>{card.desc}</p>
                    </div>
                ))}
            </div>
        </section>
    );
}
