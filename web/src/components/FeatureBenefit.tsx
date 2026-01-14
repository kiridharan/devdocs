import styles from './FeatureBenefit.module.css';

export default function FeatureBenefit() {
    const features = [
        {
            title: "CLI Integration",
            desc: "Keeps you in the terminal. No context switching required.",
            benefit: "Develop Faster",
            align: "left"
        },
        {
            title: "Multi-Language Support",
            desc: "Works with Python, Go, JS, and more. Use one tool for your whole stack.",
            benefit: "Unify Stack",
            align: "right"
        },
        {
            title: "Custom Templates",
            desc: "Match your brand. Docs that look like yours, not generic templates.",
            benefit: "Brand Identity",
            align: "left"
        }
    ];

    return (
        <section id="features" className={styles.section}>
            <div className={styles.container}>
                <div className={styles.header}>
                    <h2>Features built for developers</h2>
                </div>

                <div className={styles.list}>
                    {features.map((item, index) => (
                        <div key={index} className={`${styles.item} ${item.align === 'right' ? styles.right : ''}`}>
                            <div className={styles.text}>
                                <div className={styles.benefit}>{item.benefit}</div>
                                <h3>{item.title}</h3>
                                <p>{item.desc}</p>
                            </div>
                            <div className={styles.visual}>
                                {/* Visual placeholder */}
                                <div className={styles.placeholderBox}>Feature Visual</div>
                            </div>
                        </div>
                    ))}
                </div>
            </div>
        </section>
    );
}
