import styles from './HowItWorks.module.css';

export default function HowItWorks() {
    const steps = [
        {
            num: "01",
            title: "Connect",
            desc: "Link your repository or run our CLI tool locally. Secure by default."
        },
        {
            num: "02",
            title: "Analyze",
            desc: "DevDocs scans your functions, classes, and comments to understand your logic."
        },
        {
            num: "03",
            title: "Generate",
            desc: "Publish a stunning, searchable documentation site instantly."
        }
    ];

    return (
        <section id="how-it-works" className={styles.section}>
            <div className={styles.container}>
                <div className={styles.header}>
                    <h2>How It Works</h2>
                    <p>Three simple steps from code to docs.</p>
                </div>

                <div className={styles.steps}>
                    {steps.map((step, index) => (
                        <div key={index} className={styles.step}>
                            <div className={styles.number}>{step.num}</div>
                            <h3>{step.title}</h3>
                            <p>{step.desc}</p>
                            {index < steps.length - 1 && <div className={styles.connector} />}
                        </div>
                    ))}
                </div>
            </div>
        </section>
    );
}
