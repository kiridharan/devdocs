import styles from './FAQ.module.css';

export default function FAQ() {
    const faqs = [
        {
            q: "Who is DevDocs for?",
            a: "DevDocs is built for individual developers, agencies, and enterprise teams who want to maintain high-quality documentation without the manual effort."
        },
        {
            q: "Is it secure?",
            a: "Yes. We analyze your code locally or via secure ephemeral containers. Your source code is never stored on our servers."
        },
        {
            q: "Does it support my language?",
            a: "We currently support Python, JavaScript, TypeScript, Go, Java, and C++. More languages are coming soon."
        },
        {
            q: "Can I customize the output?",
            a: "Absolutely. You can provide custom templates and configuration files to match your brand's voice and style."
        }
    ];

    return (
        <section className={styles.section}>
            <div className={styles.container}>
                <div className={styles.header}>
                    <h2>Freqently Asked Questions</h2>
                </div>
                <div className={styles.list}>
                    {faqs.map((item, i) => (
                        <div key={i} className={styles.item}>
                            <h3 className={styles.question}>{item.q}</h3>
                            <p className={styles.answer}>{item.a}</p>
                        </div>
                    ))}
                </div>
            </div>
        </section>
    );
}
