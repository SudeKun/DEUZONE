export interface Product {
    product_id: string,
    market_id: string,
    category_id: string,
    name: string,
    keyword: string[],
    description: string,
    rating: number,
    review_count: number,
    price: number,
    quantity: number,
    images: string[],
    features: string[]
}



const products = [
    {
        product_id: "1",
        market_id: "1",
        category_id: "3",
        name: "Samsung Galaxy A70",
        keyword: ["phone", "samsung", "mobile"],
        description: "Galaxy is coming",
        rating: 4.5,
        review_count: 120,
        price: 5000,
        images: [
            "https://example.com/a70_1.jpg",
            "https://example.com/a70_2.jpg",
            "https://example.com/a70_3.jpg"
        ],
        features: [
            "6.7-inch display",
            "4500mAh battery",
            "128GB storage"
        ]
    },
    {
        product_id: "2",
        market_id: "1",
        category_id: "3",
        name: "iPhone 13",
        keyword: ["phone", "apple", "mobile"],
        description: "The latest from Apple",
        rating: 4.8,
        review_count: 230,
        price: 10999,
        images: [
            "https://example.com/iphone_13_1.jpg",
            "https://example.com/iphone_13_2.jpg",
            "https://example.com/iphone_13_3.jpg"
        ],
        features: [
            "6.1-inch display",
            "A15 Bionic chip",
            "128GB storage"
        ]
    },
    {
        product_id: "3",
        market_id: "1",
        category_id: "3",
        name: "Xiaomi Redmi Note 11",
        keyword: ["phone", "xiaomi", "mobile"],
        description: "Powerful and affordable",
        rating: 4.3,
        review_count: 150,
        price: 3500,
        images: [
            "https://example.com/redmi_note_11_1.jpg",
            "https://example.com/redmi_note_11_2.jpg",
            "https://example.com/redmi_note_11_3.jpg"
        ],
        features: [
            "6.43-inch AMOLED display",
            "5000mAh battery",
            "64GB storage"
        ]
    },
    {
        product_id: "4",
        market_id: "1",
        category_id: "3",
        name: "OnePlus 9 Pro",
        keyword: ["phone", "oneplus", "mobile"],
        description: "Fast performance and premium design",
        rating: 4.7,
        review_count: 180,
        price: 14999,
        images: [
            "https://example.com/oneplus_9_pro_1.jpg",
            "https://example.com/oneplus_9_pro_2.jpg",
            "https://example.com/oneplus_9_pro_3.jpg"
        ],
        features: [
            "6.7-inch Fluid AMOLED display",
            "4500mAh battery",
            "256GB storage"
        ]
    },
    {
        product_id: "5",
        market_id: "1",
        category_id: "3",
        name: "Huawei P40 Pro",
        keyword: ["phone", "huawei", "mobile"],
        description: "Revolutionary camera system",
        rating: 4.6,
        review_count: 100,
        price: 12000,
        images: [
            "https://example.com/huawei_p40_1.jpg",
            "https://example.com/huawei_p40_2.jpg",
            "https://example.com/huawei_p40_3.jpg"
        ],
        features: [
            "6.58-inch OLED display",
            "4200mAh battery",
            "128GB storage"
        ]
    },
    {
        product_id: "6",
        market_id: "2",
        category_id: "2",
        name: "Dell XPS 13",
        keyword: ["laptop", "dell", "ultrabook"],
        description: "Ultra-portable powerhouse",
        rating: 4.9,
        review_count: 210,
        price: 15999,
        images: [
            "https://example.com/dell_xps_13_1.jpg",
            "https://example.com/dell_xps_13_2.jpg",
            "https://example.com/dell_xps_13_3.jpg"
        ],
        features: [
            "13.4-inch 4K display",
            "Intel Core i7",
            "16GB RAM"
        ]
    },
    {
        product_id: "7",
        market_id: "2",
        category_id: "2",
        name: "MacBook Pro 16",
        keyword: ["laptop", "apple", "macbook"],
        description: "The pro laptop for professionals",
        rating: 5.0,
        review_count: 300,
        price: 24999,
        images: [
            "https://example.com/macbook_pro_16_1.jpg",
            "https://example.com/macbook_pro_16_2.jpg",
            "https://example.com/macbook_pro_16_3.jpg"
        ],
        features: [
            "16-inch Retina display",
            "Apple M1 Pro chip",
            "32GB RAM"
        ]
    },
    {
        product_id: "8",
        market_id: "2",
        category_id: "2",
        name: "HP Spectre x360",
        keyword: ["laptop", "hp", "convertible"],
        description: "Sleek and versatile laptop",
        rating: 4.6,
        review_count: 180,
        price: 13999,
        images: [
            "https://example.com/hp_spectre_1.jpg",
            "https://example.com/hp_spectre_2.jpg",
            "https://example.com/hp_spectre_3.jpg"
        ],
        features: [
            "13.3-inch OLED display",
            "Intel Core i7",
            "16GB RAM"
        ]
    },
    {
        product_id: "9",
        market_id: "2",
        category_id: "2",
        name: "Asus ZenBook 14",
        keyword: ["laptop", "asus", "ultrabook"],
        description: "Thin, light, and powerful",
        rating: 4.5,
        review_count: 170,
        price: 10999,
        images: [
            "https://example.com/asus_zenbook_1.jpg",
            "https://example.com/asus_zenbook_2.jpg",
            "https://example.com/asus_zenbook_3.jpg"
        ],
        features: [
            "14-inch Full HD display",
            "Intel Core i7",
            "16GB RAM"
        ]
    },
    {
        product_id: "10",
        market_id: "2",
        category_id: "2",
        name: "Lenovo ThinkPad X1 Carbon",
        keyword: ["laptop", "lenovo", "business"],
        description: "Designed for business professionals",
        rating: 4.7,
        review_count: 150,
        price: 16999,
        images: [
            "https://example.com/lenovo_thinkpad_1.jpg",
            "https://example.com/lenovo_thinkpad_2.jpg",
            "https://example.com/lenovo_thinkpad_3.jpg"
        ],
        features: [
            "14-inch Full HD display",
            "Intel Core i7",
            "16GB RAM"
        ]
    },
    {
        product_id: "11",
        market_id: "3",
        category_id: "4",
        name: "Sony WH-1000XM4",
        keyword: ["headphone", "sony", "noise-cancelling"],
        description: "Industry-leading noise cancellation",
        rating: 4.8,
        review_count: 220,
        price: 2499,
        images: [
            "https://example.com/sony_wh_1000xm4_1.jpg",
            "https://example.com/sony_wh_1000xm4_2.jpg",
            "https://example.com/sony_wh_1000xm4_3.jpg"
        ],
        features: [
            "Active Noise Cancelling",
            "30-hour battery life",
            "Touch controls"
        ]
    },
    {
        product_id: "12",
        market_id: "3",
        category_id: "4",
        name: "Bose QuietComfort 35 II",
        keyword: ["headphone", "bose", "noise-cancelling"],
        description: "Comfortable and clear sound",
        rating: 4.7,
        review_count: 200,
        price: 2299,
        images: [
            "https://example.com/bose_qc35_1.jpg",
            "https://example.com/bose_qc35_2.jpg",
            "https://example.com/bose_qc35_3.jpg"
        ],
        features: [
            "Noise-cancelling technology",
            "20-hour battery life",
            "Alexa-enabled"
        ]
    },
    {
        product_id: "13",
        market_id: "3",
        category_id: "4",
        name: "JBL Live 650BTNC",
        keyword: ["headphone", "jbl", "bluetooth"],
        description: "Great sound and bass",
        rating: 4.4,
        review_count: 120,
        price: 799,
        images: [
            "https://example.com/jbl_live_650_1.jpg",
            "https://example.com/jbl_live_650_2.jpg",
            "https://example.com/jbl_live_650_3.jpg"
        ],
        features: [
            "Active Noise Cancelling",
            "30-hour battery life",
            "Bluetooth 5.0"
        ]
    },
    {
        product_id: "14",
        market_id: "3",
        category_id: "4",
        name: "Sennheiser Momentum 3",
        keyword: ["headphone", "sennheiser", "premium"],
        description: "Premium sound quality",
        rating: 4.9,
        review_count: 250,
        price: 3699,
        images: [
            "https://example.com/sennheiser_momentum_1.jpg",
            "https://example.com/sennheiser_momentum_2.jpg",
            "https://example.com/sennheiser_momentum_3.jpg"
        ],
        features: [
            "Noise-cancelling",
            "17-hour battery life",
            "Wireless"
        ]
    },
    {
        product_id: "15",
        market_id: "3",
        category_id: "4",
        name: "Beats Studio 3 Wireless",
        keyword: ["headphone", "beats", "wireless"],
        description: "Sound that moves you",
        rating: 4.6,
        review_count: 190,
        price: 2999,
        images: [
            "https://example.com/beats_studio_3_1.jpg",
            "https://example.com/beats_studio_3_2.jpg",
            "https://example.com/beats_studio_3_3.jpg"
        ],
        features: [
            "Noise-cancelling",
            "22-hour battery life",
            "Bluetooth connectivity"
        ]
    },

];

export default products;
